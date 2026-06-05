# Wine Island

**Équipe :** Alexis BAHUAUD, Paul DECATOIRE, Ronan DUPAS, Axel LOQUET, Godwin OBLASSE, Tayvadi PHAISAN, Emmanuel YOHORE. 
**Groupe / promo :** Wine Island / INFO B3  
**Dépôt :** https://github.com/CODEBYGODWIN/DevOpsProject.git 

---

## Description du sujet

L'application est un site vitrine permettant de présenter des projets, 
des compétences et des informations de contact. L'utilisateur cible est 
toute personne souhaitant consulter un portfolio en ligne, comme un 
recruteur ou un client potentiel. Le backend exposera des endpoints REST 
en Go (Gin) pour servir les données stockées dans PostgreSQL, tandis que 
le front Vue.js affichera les contenus de manière dynamique. Le projet 
servira également de support d'apprentissage DevOps : containerisation, 
pipeline CI et déploiement Kubernetes.

---

## Stack technique prévu

| Composant | Choix | Justification (1 phrase) |
| --------- | ----- | -------------------------- |
| Backend / API | Go (Framework chi) | Go et chi permettent de créer une API REST légère, rapide à builder et simple à conteneuriser. |
| Base de données | PostgreSQL | PostgreSQL est une base relationnelle fiable, adaptée aux données applicatives et facile à lancer en service Docker. |
| Front (optionnel) | Vue.js | Vue.js permet de construire une interface dynamique tout en gardant une structure front simple pour le projet. |
| Orchestration cible | Compose puis K8s | Docker Compose assure un démarrage reproductible en local, avec une ouverture possible vers Kubernetes si le périmètre évolue. |

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
| S2 | Dockerfile(s) + DB en container | ☐ |
| S3 | docker-compose + CI vert | ☐ |
| S4 | Manifests K8s appliqués | ☐ |
| S5 | Monitoring + post-mortem | ☐ |
| S6 | Soutenance prête | ☐ |

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

---

## Participation S1 (optionnel, 2 lignes)

Retour sur le jeu de rôle ou le cas déploiement : une leçon retenue pour le projet.
