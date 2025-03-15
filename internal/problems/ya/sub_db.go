package ya

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"sort"
)

// SubscribeDatabase задача про реализацию простой базы данных с возможностью подписки на изменение некоторых полей сущности.
func SubscribeDatabase(reader io.Reader, writer io.Writer) {
	r := bufio.NewReader(reader)
	w := bufio.NewWriter(writer)

	var (
		subC int
		reqC int
	)

	_, _ = fmt.Fscanf(reader, "%d %d\n", &subC, &reqC)

	db := &Database{
		Bus: &Bus{
			Reader: r,
			Writer: w,
		},
		Subs:      make([]*Subscriber, 0, 50),
		Offers:    make(map[string]*Offer, 10000),
		ObsFields: make(map[string]*ObservableField, 0),
	}

	for ; subC > 0; subC-- {
		db.AcceptSubscriber()
	}

	db.Run(reqC)

	_ = w.Flush()
}

// Offer stores all offer's data.
type Offer struct {
	ID             string
	Price          float64
	StockCount     float64
	PartnerContent struct {
		Title       string
		Description string
	}
}

func (offer *Offer) Update(fName string, value any) {
	switch fName {
	case "id":
		offer.ID = value.(string)
	case "price":
		offer.Price = value.(float64)
	case "stock_count":
		offer.StockCount = value.(float64)
	case "title":
		offer.PartnerContent.Title = value.(string)
	case "description":
		offer.PartnerContent.Description = value.(string)
	default:
		panic(fmt.Sprintf("unsupported field to update: %s", fName))
	}
}

// GetFieldValue returns level 1 representation of field's value.
// TODO: God, bless this trash code below
func (offer *Offer) GetFieldValue(fName string) (any, bool) {
	switch fName {
	case "id":
		if offer.ID == "" {
			return nil, false
		}
		return offer.ID, true
	case "price":
		if offer.Price == 0 {
			return nil, false
		}
		return offer.Price, true
	case "stock_count":
		if offer.StockCount == 0 {
			return nil, false
		}
		return offer.StockCount, true
	case "title":
		if offer.PartnerContent.Title == "" {
			return nil, false
		}
		return offer.PartnerContent.Title, true
	case "description":
		if offer.PartnerContent.Description == "" {
			return nil, false
		}
		return offer.PartnerContent.Description, true
	case "partner_content":
		switch {
		case offer.PartnerContent.Title != "" && offer.PartnerContent.Description != "":
			return map[string]interface{}{
				"title":       offer.PartnerContent.Title,
				"description": offer.PartnerContent.Description,
			}, true
		case offer.PartnerContent.Title != "":
			return map[string]interface{}{
				"title": offer.PartnerContent.Title,
			}, true
		case offer.PartnerContent.Description != "":
			return map[string]interface{}{
				"description": offer.PartnerContent.Description,
			}, true
		default:
			return nil, false
		}
	default:
		panic(fmt.Sprintf("unsupported field to update: %s", fName))
	}
}

func (offer *Offer) ToPart(fields ...string) OfferPart {
	part := OfferPart{
		"id":              offer.ID,
		"partner_content": make(map[string]interface{}, 2),
	}

	for _, field := range fields {
		switch field {
		case "title", "description":
			if val, ok := offer.GetFieldValue(field); ok {
				part["partner_content"].(map[string]interface{})[field] = val
			}
		default:
			if val, ok := offer.GetFieldValue(field); ok {
				part[field] = val
			}
		}
	}

	if len(part["partner_content"].(map[string]interface{})) == 0 {
		delete(part, "partner_content")
	}

	return part
}

type OfferPart map[string]interface{}

// Message is an update and response body.
type Message struct {
	TraceID string    `json:"trace_id"`
	Offer   OfferPart `json:"offer"`
}

// Subscriber holds subscriber descriptor.
type Subscriber struct {
	ID int

	Triggers      []string
	ProvideFields []string // contains trigger and shipment fields together
}

// Database должен быть собран из списка запросов на подписку и отправлять события на обновление данных.
type Database struct {
	Bus       *Bus
	Offers    map[string]*Offer
	Subs      []*Subscriber
	ObsFields map[string]*ObservableField // Access by field
}

// AcceptSubscriber accepts subscribe requests.
func (d *Database) AcceptSubscriber() {
	sub := d.Bus.ReadSubscriber()
	d.Subs = append(d.Subs, sub)

	// update observable fields
	for _, trigger := range sub.Triggers {
		obs, exists := d.ObsFields[trigger]
		if !exists {
			d.ObsFields[trigger] = &ObservableField{
				Field: trigger,
				Subs:  []*Subscriber{sub},
			}
		} else {
			obs.Subs = append(obs.Subs, sub)
		}
	}
}

func (d *Database) Run(requestsCount int) {
	for ; requestsCount > 0; requestsCount-- {
		d.handle(d.Bus.GetUpdateRequest())
	}
}

func (d *Database) handle(req *Message) {
	subSet := make(map[int]*Subscriber, 0)
	addFieldSubscribers := func(f string) {
		if obs, exists := d.ObsFields[f]; exists {
			for _, sub := range obs.Subs {
				subSet[sub.ID] = sub
			}
		}
	}

	id := req.Offer["id"].(string)
	if _, exists := d.Offers[id]; !exists {
		d.Offers[id] = &Offer{ID: id}
	}

	for field, value := range req.Offer {
		// TODO: better solution ?
		if field == "partner_content" {
			addFieldSubscribers(field)

			for k, v := range value.(map[string]interface{}) {
				addFieldSubscribers(k)
				d.Offers[id].Update(k, v)
			}
		} else {
			// common level 1 field
			addFieldSubscribers(field)
			d.Offers[id].Update(field, value)
		}
	}

	subs := make([]*Subscriber, 0, len(subSet))
	for _, s := range subSet {
		subs = append(subs, s)
	}

	d.notify(d.Offers[id], subs, req.TraceID)
}

func (d *Database) notify(offer *Offer, subs []*Subscriber, traceID string) {
	sort.Slice(subs, func(i, j int) bool {
		return subs[i].ID < subs[j].ID
	})

	for _, sub := range subs {
		msg := &Message{
			TraceID: traceID,
		}

		msg.Offer = offer.ToPart(sub.ProvideFields...)

		d.Bus.Notify(msg)
	}
}

// ObservableField describes a single field of offer,
// that subscribers want to receive updates of.
type ObservableField struct {
	Field string
	Subs  []*Subscriber
}

type Bus struct {
	Reader *bufio.Reader
	Writer *bufio.Writer
	subSeq int // [0, 49]
}

func (b *Bus) ReadSubscriber() *Subscriber {
	triggerC, shipmentC := 0, 0
	_, _ = fmt.Fscanf(b.Reader, "%d %d", &triggerC, &shipmentC)

	var sub = Subscriber{
		ID:            b.subSeq,
		Triggers:      make([]string, 0, triggerC),
		ProvideFields: make([]string, 0, triggerC+shipmentC),
	}
	b.subSeq++

	var field string
	for i := triggerC; i > 0; i-- {
		_, _ = fmt.Fscanf(b.Reader, "%s", &field)

		sub.Triggers = append(sub.Triggers, field)
	}

	sub.ProvideFields = append(sub.ProvideFields, sub.Triggers...)
	for i := shipmentC; i > 0; i-- {
		_, _ = fmt.Fscanf(b.Reader, "%s", &field)

		sub.ProvideFields = append(sub.ProvideFields, field)
	}

	_, _ = fmt.Fscanf(b.Reader, "\n")

	return &sub
}

func (b *Bus) GetUpdateRequest() *Message {
	// By task definition, errors in json read and parsing are impossible.
	raw, _ := b.Reader.ReadBytes('\n')

	var m Message
	_ = json.Unmarshal(raw, &m)

	return &m
}

func (b *Bus) Notify(msg *Message) {
	d, _ := json.Marshal(msg)
	_, _ = fmt.Fprintln(b.Writer, string(d))
}
