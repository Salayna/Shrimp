package main

/*
TODO: Créer les flags (-h : afiche l'aide
 -f donne un fichier de configuration
 -i passe en mode interactif
  si pas d'option affiche le -h
-v affiche en mode verbose toute les info de création
-l language
-p Nom du Projet)
TODO: créer un dossier à la racine à partir du nom (parse également )
TODO: Parser les fichiers yaml (
	fichier de configuration principal de projet
	fichier de language dans ./config

)
-TODO: Mettre à Jour le README à la fin de la session
*/
import (
	"github.com/Salayna/create-project-cli/internal/cli"
)

func main() {
	cli.Cli()
}
