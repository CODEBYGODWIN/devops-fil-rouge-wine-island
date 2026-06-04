# DevOpsProject

## Conteneurisation de l'API

L'API backend est une application Go situee dans `backend/`. Le Dockerfile utilise une image officielle `golang:1.25.3-alpine3.22` pour compiler le binaire, puis une image finale `alpine:3.22` afin de garder une image legere avec un tag explicite.

```bash
docker build -t mon-api:local .
docker run --rm -p 3000:3000 mon-api:local
curl http://localhost:3000/health
```

Les variables sensibles de base de donnees (`DB_HOST`, `DB_USER`, `DB_PASSWORD`, etc.) ne sont pas integrees a l'image. Elles doivent etre passees au lancement avec `-e` ou via un fichier d'environnement local.
