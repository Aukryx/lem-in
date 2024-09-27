#!/bin/bash

# Script : build_animation.sh
# Description : Exécute la simulation, convertit les fichiers DOT en PNG, crée une animation,
#               et supprime les fichiers temporaires en conservant seulement le GIF et la vidéo.

# Fonction pour afficher un message d'erreur et quitter
function error_exit {
    echo "$1" 1>&2
    exit 1
}

# Vérifier si le programme visualizer existe
if ! [ -x "./visualizer" ]; then
  error_exit "Erreur : Le programme 'visualizer' n'existe pas ou n'est pas exécutable."
fi

# Vérifier si le fichier d'entrée existe
if ! [ -f "../examples/example.txt" ]; then
  error_exit "Erreur : Le fichier d'entrée '../example/example00.txt' n'existe pas."
fi

echo "1. Exécution de la simulation..."
./visualizer ../examples/example07.txt || error_exit "Erreur lors de l'exécution de la simulation."

echo "2. Conversion des fichiers DOT en PNG..."
for i in step_*.dot; do
    if [ -f "$i" ]; then
        dot -Tpng "$i" -o "${i%.dot}.png" || error_exit "Erreur lors de la conversion de $i en PNG."
    else
        echo "Aucun fichier $i trouvé."
    fi
done

echo "3. Création de l'animation avec ffmpeg..."
ffmpeg -framerate 1 -i step_%d.png -vf "scale=trunc(iw/2)*2:trunc(ih/2)*2" -c:v libx264 -pix_fmt yuv420p animation.mp4
if [ $? -ne 0 ]; then
    error_exit "Erreur lors de la création de la vidéo avec ffmpeg."
fi

echo "4. Création de l'animation GIF avec ImageMagick..."
convert -delay 100 step_*.png animation.gif
if [ $? -ne 0 ]; then
    error_exit "Erreur lors de la création du GIF animé avec ImageMagick."
fi

echo "5. Suppression des fichiers temporaires (.dot et .png)..."
rm step_*.dot step_*.png
if [ $? -ne 0 ]; then
    error_exit "Erreur lors de la suppression des fichiers temporaires."
fi

echo "Animation créée avec succès :"
echo " - Vidéo : animation.mp4"
echo " - GIF animé : animation.gif"
