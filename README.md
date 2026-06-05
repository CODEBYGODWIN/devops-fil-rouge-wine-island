# Wine Island

**Équipe :**  Alexis BAHUAUD, Paul DECATOIRE, Ronan DUPAS, Axel LOQUET, Godwin OBLASSE, Tayvadi PHAISAN, Emmanuel YOHORE.

**Groupe / promo :** Wine Island / INFO B3

**Dépôt :** https://github.com/CODEBYGODWIN/DevOpsProject.git 

---

## Description du sujet

L'application Wine Island est une solution de gestion dédiée au suivi d'une 
sélection de vins, de domaines et d'informations associées. Elle permet de 
centraliser les données, de les consulter et de les administrer depuis une 
interface web. L'utilisateur cible est une équipe souhaitant gérer un catalogue 
de vins de manière simple et structurée. Le backend exposera des endpoints REST 
en Go avec chi pour servir les données stockées dans PostgreSQL, tandis que le 
front Vue.js affichera les contenus de manière dynamique. Le projet servira 
également de support d'apprentissage DevOps : conteneurisation, Docker Compose 
et pipeline CI/CD.

---

## Stack technique prévu

| Composant | Choix | Justification (1 phrase) |
| --------- | ----- | -------------------------- |
| Backend / API | Go (Framework chi) | Go et chi permettent de créer une API REST légère, rapide à builder et simple à conteneuriser. |
| Base de données | PostgreSQL | PostgreSQL est une base relationnelle fiable, adaptée aux données applicatives et facile à lancer en service Docker. |
| Front (optionnel) | Vue.js | Vue.js permet de construire une interface dynamique tout en gardant une structure front simple pour le projet. |
| Orchestration locale | Docker Compose | Docker Compose assure un démarrage reproductible en local avec les services API et base de données. |
| CI/CD | GitHub Actions ou GitLab CI | La CI/CD permet de vérifier le code, lancer les tests, construire l'image Docker et la publier automatiquement sur un registre. |

---

## Rôles dans l'équipe

| Membre | Rôle | Responsabilité principale |
| ------ | ---- | ------------------------- |
| Emmanuel YOHORE & Axel LOQUET | Lead Dev | Développer l'API, structurer le code applicatif et vérifier que les endpoints répondent aux besoins du sujet. |
| Godwin OBLASSE & Paul DECATOIRE | Lead Ops | Préparer les Dockerfiles, le docker-compose, les variables d'environnement et le diagnostic des services. |
| Tayvadi PHAISAN & Alexis BAHUAUD | Lead Qualité / CI | Mettre en place le lint, les tests, le build d'image et le push registry dans la pipeline CI/CD. |
| Ronan DUPAS | Lead Doc / Produit | Maintenir le README, suivre les jalons, clarifier le besoin produit et préparer la trame de soutenance. |

---

## Objectifs du fil rouge (3 minimum)

1. Développer une API REST en Go pour gérer les contenus du site vitrine et les stocker dans PostgreSQL.
2. Conteneuriser l'application avec au moins deux services reproductibles : API et base de données.
3. Mettre en place une chaîne CI/CD qui vérifie le code, lance les tests, construit l'image Docker et la publie sur un registre.

---

## Jalons — état d'avancement

| Séance | Livrable | Statut (à cocher) |
| ------ | -------- | ----------------- |
| S1 | README cadrage | ☑ |
| S2 | Dockerfile(s) + DB en container | ☑ |
| S3 | docker-compose + CI vert | ☑ |
| S4 | Manifests K8s appliqués | ☐ |
| S5 | Monitoring + post-mortem | ☐ |
| S6 | Soutenance prête | ☑ |

---

## Démarrage local (à compléter au fil des séances)

```bash
git clone https://github.com/CODEBYGODWIN/DevOpsProject.git
cd DevOpsProject
docker compose up -d
```

---

## Communication d'équipe

Canal utilisé : Discord
