# Go Article API

Ceci est une application Web simple en Go (golang) qui expose une API pour gérer des articles. L'API permet de récupérer tous les articles existants et d'ajouter de nouveaux articles.

## Installation

1. Assurez-vous d'avoir Go (golang) installé sur votre machine.
2. Clonez ce dépôt ou téléchargez les fichiers source.

## Utilisation

1. Accédez au répertoire du projet.

2. Exécutez la commande suivante pour installer les dépendances :

go mod download


3. Lancez l'application en utilisant la commande :

go run main.go


4. L'application sera maintenant accessible à l'adresse `http://localhost:8080`.

## Endpoints

- `GET /` : Renvoie un message de bienvenue sur la page d'accueil.

- `GET /articles` : Renvoie tous les articles existants au format JSON.

- `POST /articles` : Permet de créer un nouvel article. L'article doit être envoyé au format JSON dans le corps de la requête.

## Exemple d'article

Voici un exemple d'article au format JSON :

```json
{
"Title": "Mon titre",
"Desc": "Ma description",
"Content": "Mon contenu"
}
```

Assurez-vous d'inclure les champs Title, Desc et Content lors de la création d'un nouvel article.

N'hésitez pas à personnaliser ce fichier README en fonction des besoins de votre projet.
