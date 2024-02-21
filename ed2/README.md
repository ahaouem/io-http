# Komendy curl do uzycia, aby przetestować endpointy

## Pobieranie wszystkich uzytkownikow

```bash
curl http://127.0.0.1:5000/users
```

## Pobieranie użytkownika po ID

```bash
curl http://127.0.0.1:5000/users/1
```

## Dodawanie nowego użytkownika

```bash
curl -X POST -H "Content-Type: application/json" -d '{"firstName": "aleks", "lastName": "hauauauaem", "birth_year": 1994, "group": "admin"}' http://127.0.0.1:5000/users

```

## Aktualizacja danych użytkownika

```bash
 curl -X PATCH http://127.0.0.1:5000/users/1 \
-H "Content-Type: application/json" \
-d '{"firstName": "UpdatedName", "lastName": "UpdatedLastName", "birthYear": 1995, "group": "premium"}'
```

## Usuwanie użytkownika

```bash
curl -X DELETE http://127.0.0.1:5000/users/1
```
