# Gabarit README — Projet fil rouge DevOps

> Copiez ce contenu dans le `README.md` à la racine du **dépôt Git de votre groupe**.  
> Complétez chaque section — champs vides = Note 3 incomplète.

---

# [Nom du projet]

**Équipe :** Paul DECATOIRE, Godwin OBLASSE, Emmanuel YOHORE, Tayvadi PHAISAN, Ronan DUPAS. 
**Groupe / promo :** B3  
**Dépôt :** https://github.com/CODEBYGODWIN/DevOpsProject.git

---

## Description du sujet

En 3–5 phrases : que fait l'application ? Qui est l'utilisateur cible ?

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
| | Lead Dev | |
| | Lead Ops | |
| | Lead Qualité / CI | |
| | Lead Doc / Produit | |

---

## Objectifs du fil rouge (3 minimum)

1. Ex. : Avoir une API conteneurisée avec healthcheck d'ici S3.
2. Ex. : Pipeline CI qui build et push l'image sur chaque merge `main`.
3. Ex. : Déployer sur cluster kind avec 2 replicas d'ici S4.

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
# À documenter progressivement — pas besoin de tout remplir en S1
git clone ...
```

---

## Communication d'équipe

Canal utilisé (Teams, Discord, …) :

---

## Participation S1 (optionnel, 2 lignes)

Retour sur le jeu de rôle ou le cas déploiement : une leçon retenue pour le projet.