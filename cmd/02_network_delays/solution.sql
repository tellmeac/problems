select (select sum(@(extract(epoch from child.datetime) * 1000 - extract(epoch from parent.datetime) * 1000)) as delay
        from requests child
                 join requests parent on (child.parent_request_id = parent.request_id)
        where (child.host = substring(parent.data from '\S*'))
          and (
                (child.type = 'RequestReceived' and parent.type = 'RequestSent') or
                (child.type = 'ResponseSent' and parent.type = 'ResponseReceived')
            )) /
       (select count(request_id)
        from requests
        where host = 'balancer.test.yandex.ru'
          and type = 'RequestReceived') as avg_network_time_ms;