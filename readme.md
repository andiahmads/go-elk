# GO-ELK v1.0.8 🚀

Implement Forwarding log to ELK with monitoring using KIBANA
- using web hook


# Setup & Running Kibana & ELK:
```env
docker-compose up -d
```

# kibana dashboard
```env
http://localhost:5601/
```

# run:
```env
go run main.go
```


# API Specifiation:
```json
{
  "endpoint": "localhost:6060",
  "method": "GET",
}
```


# testing:
```env
curl --location --request GET locahost:6060/
```