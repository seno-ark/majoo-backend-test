# Majoo Backend Test

## How to run
```
make run
```

___

## API Endpoints

### Login

**POST /login**

Form Data:
- username
- password

Request
```
curl -XPOST 'localhost:9000/login' -d 'username=admin1' -d 'password=admin1'
```

Response
```
{
  "message": "Success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcl9uYW1lIjoiYWRtaW4xIiwiZXhwIjoxNjQ4MzkyMjk2fQ.EkEzYvvPu9NbpxeMC6PvJvFuWYuJ27-XlhN93rUo8To",
    "id": 1,
    "name": "Admin 1",
    "user_name": "admin1"
  }
}
```

### Merchant Report

**GET /report/merchant/:merchant_id/omzet**

Query Params:
- `start_date` (YYYY-MM-DD) default: today's date
- `end_date` (YYYY-MM-DD) default: today's date
- `page` default: 1
- `count` default: 20

Request
```
curl -XGET 'localhost:9000/report/merchant/1/omzet?start_date=2021-11-01&end_date=2021-11-30&page=1&count=5' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcl9uYW1lIjoiYWRtaW4xIiwiZXhwIjoxNjQ4MzkyMjk2fQ.EkEzYvvPu9NbpxeMC6PvJvFuWYuJ27-XlhN93rUo8To'
```

Response
```
{
  "message": "",
  "data": {
    "merchant_id": 1,
    "outlet_id": 0,
    "merchant_name": "merchant 1",
    "outlet_name": "",
    "omzets": [
      {
        "date": "2021-11-01T00:00:00Z",
        "total": 4500
      },
      {
        "date": "2021-11-02T00:00:00Z",
        "total": 6000
      },
      {
        "date": "2021-11-03T00:00:00Z",
        "total": 2500
      },
      {
        "date": "2021-11-04T00:00:00Z",
        "total": 6000
      },
      {
        "date": "2021-11-05T00:00:00Z",
        "total": 14000
      }
    ],
    "total_data": 30,
    "filters": {
      "count": 5,
      "page": 1
    }
  }
}
```

### Outlet Report

**GET /report/outlet/:outlet_id/omzet**

Query Params:
- `start_date` (YYYY-MM-DD) default: today's date
- `end_date` (YYYY-MM-DD) default: today's date
- `page` default: 1
- `count` default: 20

Request
```
curl -XGET 'localhost:9000/report/outlet/3/omzet?start_date=2021-11-01&end_date=2021-11-20&page=2&count=3' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcl9uYW1lIjoiYWRtaW4xIiwiZXhwIjoxNjQ4MzkyMjk2fQ.EkEzYvvPu9NbpxeMC6PvJvFuWYuJ27-XlhN93rUo8To'
```

Response
```
{
  "message": "",
  "data": {
    "merchant_id": 1,
    "outlet_id": 3,
    "merchant_name": "merchant 1",
    "outlet_name": "Outlet 2",
    "omzets": [
      {
        "date": "2021-11-04T00:00:00Z",
        "total": 5000
      },
      {
        "date": "2021-11-05T00:00:00Z",
        "total": 7000
      },
      {
        "date": "2021-11-06T00:00:00Z",
          "total": 0
      }
    ],
    "total_data": 20,
    "filters": {
      "count": 3,
      "page": 2
    }
  }
}
```
