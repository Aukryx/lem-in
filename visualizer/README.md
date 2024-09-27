
# README

## **Lem-in : Simulation et Visualisation des Mouvements de Fourmis**

Ce projet est une implémentation en Go d'une simulation de mouvements de fourmis à travers un réseau de salles et de tunnels. Le programme lit une description du réseau (salles, tunnels, fourmis) à partir d'un fichier, calcule les meilleurs chemins pour les fourmis, et génère des fichiers de visualisation pour représenter les mouvements étape par étape.

---

## **Table des Matières**

- [Prérequis](#prérequis)
- [Structure du Projet](#structure-du-projet)
- [Compilation et Exécution du Programme](#compilation-et-exécution-du-programme)
- [Génération des Fichiers DOT](#génération-des-fichiers-dot)
- [Conversion des Fichiers DOT en Images](#conversion-des-fichiers-dot-en-images)
- [Création de l'Animation](#création-de-lanimation)
- [Exemple Complet](#exemple-complet)
- [Remarques et Conseils](#remarques-et-conseils)

---

## **Prérequis**

Avant de commencer, assurez-vous d'avoir installé les logiciels suivants sur votre système :

1. **Go** (version 1.16 ou supérieure) : [Installer Go](https://golang.org/dl/)
2. **Graphviz** : Pour convertir les fichiers DOT en images.
   - Installation sous Ubuntu/Debian :
     ```bash
     sudo apt-get install graphviz
     ```
   - Installation sous macOS avec Homebrew :
     ```bash
     brew install graphviz
     ```
3. **ffmpeg** ou **ImageMagick** : Pour créer l'animation à partir des images.
   - **ffmpeg** :
     - Installation sous Ubuntu/Debian :
       ```bash
       sudo apt-get install ffmpeg
       ```
     - Installation sous macOS avec Homebrew :
       ```bash
       brew install ffmpeg
       ```
   - **ImageMagick** :
     - Installation sous Ubuntu/Debian :
       ```bash
       sudo apt-get install imagemagick
       ```
     - Installation sous macOS avec Homebrew :
       ```bash
       brew install imagemagick
       ```

---

## **Structure du Projet**

- **main.go** : Le fichier source principal contenant le code du programme.
- **example.txt** : Un fichier d'entrée exemple décrivant le réseau de salles et de tunnels.
- **step_*.dot** : Fichiers DOT générés par le programme pour chaque étape des mouvements des fourmis.
- **step_*.png** : Images générées à partir des fichiers DOT.
- **animation.mp4** ou **animation.gif** : Animation des mouvements des fourmis.

---

## **Compilation et Exécution du Programme**

### **1. Compilation du Programme**

Ouvrez un terminal dans le répertoire du projet et compilez le programme Go :

```bash
go build main.go
```

Cela générera un exécutable nommé `main` (ou `main.exe` sous Windows).

### **2. Préparation du Fichier d'Entrée**

Le fichier d'entrée doit respecter le format suivant :

- La première ligne indique le **nombre de fourmis**.
- Les salles sont définies avec leur nom et leurs coordonnées X et Y :
  ```
  SalleX X Y
  ```
- Les liens entre les salles sont définis par :
  ```
  SalleX-SalleY
  ```
- Les directives `##start` et `##end` indiquent les salles de départ et d'arrivée.

**Exemple (`example.txt`) :**

```
4
##start
A 1 1
B 5 1
C 3 3
D 1 5
E 5 5
##end
F 3 7
A-B
A-D
B-C
C-E
D-E
E-F
```

### **3. Exécution du Programme**

Exécutez le programme en spécifiant le fichier d'entrée :

```bash
./main example.txt
```

Le programme va :

- Lire et analyser le fichier d'entrée.
- Calculer les meilleurs chemins pour les fourmis.
- Générer des fichiers DOT pour chaque étape des mouvements des fourmis (`step_0.dot`, `step_1.dot`, etc.).
- Afficher les informations et les mouvements dans le terminal.

---

## **Génération des Fichiers DOT**

Les fichiers DOT (`step_0.dot`, `step_1.dot`, ...) sont générés automatiquement par le programme lors de son exécution. Chaque fichier représente l'état du réseau de salles et les positions des fourmis à un instant donné.

---

## **Conversion des Fichiers DOT en Images**

Pour visualiser les graphes, vous devez convertir les fichiers DOT en images (PNG).

### **Commande pour Convertir un Fichier DOT en PNG**

```bash
dot -Tpng step_0.dot -o step_0.png
```

### **Automatiser la Conversion pour Tous les Fichiers DOT**

Vous pouvez utiliser une boucle pour convertir tous les fichiers DOT générés :

```bash
for i in step_*.dot; do
    dot -Tpng "$i" -o "${i%.dot}.png"
done
```

---

## **Création de l'Animation**

### **Option 1 : Utiliser ffmpeg pour Créer une Vidéo**

```bash
ffmpeg -framerate 1 -i step_%d.png -vf "scale=trunc(iw/2)*2:trunc(ih/2)*2" -c:v libx264 -pix_fmt yuv420p animation.mp4
```

- **-framerate 1** : Définit le nombre d'images par seconde (ajustez la valeur selon vos préférences).
- **-i step_%d.png** : Indique que les images d'entrée sont nommées `step_0.png`, `step_1.png`, etc.
- **-vf "scale=trunc(iw/2)*2:trunc(ih/2)*2"** : Redimensionne les images pour que les dimensions soient divisibles par 2.
- **animation.mp4** : Nom du fichier vidéo de sortie.

### **Option 2 : Utiliser ImageMagick pour Créer un GIF Animé**

```bash
convert -delay 100 step_*.png animation.gif
```

- **-delay 100** : Définit le délai entre les images en centièmes de seconde.
- **animation.gif** : Nom du fichier GIF animé de sortie.

---

## **Exemple Complet**

### **1. Exécution du Programme**

```bash
./visualizer ../example/example.txt
```

### **2. Conversion des Fichiers DOT en Images**

```bash
for i in step_*.dot; do
    dot -Tpng "$i" -o "${i%.dot}.png"
done
```

### **3. Création de l'Animation**

**Avec ffmpeg :**

```bash
ffmpeg -framerate 1 -i step_%d.png -vf "scale=trunc(iw/2)*2:trunc(ih/2)*2" -c:v libx264 -pix_fmt yuv420p animation.mp4
```

**Avec ImageMagick :**

```bash
convert -delay 100 step_*.png animation.gif
```

### **4. Visualisation de l'Animation**

- **Pour la vidéo :** Ouvrez `animation.mp4` avec votre lecteur vidéo préféré.
- **Pour le GIF animé :** Ouvrez `animation.gif` avec votre navigateur ou une visionneuse d'images.

---

## **Remarques et Conseils**

### **Personnalisation des Graphes**

- **Couleurs et Styles :** Vous pouvez modifier les attributs dans la fonction `generateDOTFile` pour changer les couleurs des nœuds, la forme, les styles des arêtes, etc.
- **Échelle des Coordonnées :** Si les nœuds sont trop espacés ou trop rapprochés, ajustez la variable `scale` dans `generateDOTFile`.
  ```go
  scale := 1.0 // Vous pouvez ajuster cette valeur
  ```

### **Gestion des Dimensions**

- Si vous rencontrez des problèmes avec des images coupées, assurez-vous que les attributs `size`, `ratio`, et `dpi` sont correctement définis dans `generateDOTFile`.
- Vous pouvez également ajuster ces attributs lors de la conversion avec `dot`.

### **Dépendances**

- Assurez-vous que les outils **Graphviz**, **ffmpeg**, et **ImageMagick** sont correctement installés et accessibles depuis le terminal.
- Sous Windows, vous pouvez utiliser des environnements comme **WSL (Windows Subsystem for Linux)** ou installer les versions Windows de ces outils.

---

## **Contact**

Pour toute question ou assistance supplémentaire, n'hésitez pas à me contacter ou à consulter la documentation des outils utilisés.

---

**Bon courage et bonne visualisation !**
