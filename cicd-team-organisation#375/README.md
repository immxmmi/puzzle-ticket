


## Admin-Benutzer erstellen

GlitchTip hat keinen Standard-Login. Du musst selbst einen Benutzer anlegen:

```bash
docker compose up -d
docker compose exec glitchtip python manage.py createsuperuser
```

Danach kannst du dich unter [http://localhost:3000](http://localhost:3000) mit den angegebenen Zugangsdaten anmelden.