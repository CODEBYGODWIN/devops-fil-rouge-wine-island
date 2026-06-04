# Wine Island

**Équipe :** Paul DECATOIRE, Godwin OBLASSE, Emmanuel YOHORE, Tayvadi PHAISAN, Ronan DUPAS. 
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
| Backend / API | Go (Framework chi) | |
| Base de données | PostgreSQL | |
| Front (optionnel) | Vue.js | |
| Orchestration cible | Compose puis K8s | |

---

## Rôles dans l'équipe

| Membre | Rôle | Responsabilité principale |
| ------ | ---- | ------------------------- |
| Emmanuel YOHORE & Axel LOQUET | Lead Dev | |
| Godwin OBLASSE & Paul DECATOIRE | Lead Ops | |
| Tayvadi PHAISAN & Alexis BAHUAUD | Lead Qualité / CI | |
| Ronan DUPAS | Lead Doc / Produit | |

---

## Objectifs du fil rouge (3 minimum)

1. Avoir une API conteneurisée avec healthcheck d'ici S3.
2. Pipeline CI qui build et push l'image sur chaque merge `main`.
3. Déployer sur cluster kind avec 2 replicas d'ici S4.

---

## Jalons — état d'avancement

| Séance | Livrable | Statut (à cocher) |
| ------ | -------- | ----------------- |
| S1 | README cadrage | ☐ |
| S2 | Dockerfile(s) + DB en container | ☐ |
| S3 | docker-compose + CI vert | ☐ |
| S4 | Manifests K8s appliqués | ☐ |
| S5 | Monitoring + post-mortem | ☐ |
| S6 | Soutenance prête | ☐ |

---

## Démarrage local (à compléter au fil des séances)

```bash
git clone ...
```

---

## Communication d'équipe

Canal utilisé : Discord

---

## Participation S1 (optionnel, 2 lignes)

Retour sur le jeu de rôle ou le cas déploiement : une leçon retenue pour le projet.