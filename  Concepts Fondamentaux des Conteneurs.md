# Concepts Fondamentaux des Conteneurs Linux

Ce document explique les trois concepts clés qui constituent la base de la technologie des conteneurs Linux.

## Namespaces

Les namespaces sont un mécanisme d'isolation qui permet de limiter ce qu'un processus peut voir du système. Ils créent une abstraction qui fait croire au processus qu'il a son propre environnement isolé.

### Types principaux de namespaces :
- **PID** : Isole les identifiants de processus
- **Network** : Isole les interfaces réseau
- **Mount** : Isole les points de montage du système de fichiers
- **UTS** : Isole les noms d'hôte
- **IPC** : Isole les ressources de communication inter-processus
- **User** : Isole les identifiants d'utilisateurs et de groupes

Les namespaces permettent à un conteneur de fonctionner comme s'il était un système indépendant, sans voir les processus, réseaux ou utilisateurs du système hôte.

## Chroot

Chroot (Change Root) est une opération qui change le répertoire racine apparent pour un processus et ses enfants. C'est l'une des premières formes d'isolation sur Unix/Linux.

### Fonctionnement :
- Un processus "chrooté" ne peut pas accéder aux fichiers en dehors de son répertoire racine modifié
- Il crée une cage de fichiers (jail) qui limite ce que le processus peut voir du système de fichiers
- C'est un mécanisme simple mais limité car il n'isole que le système de fichiers, pas les autres ressources

Chroot est un ancêtre des conteneurs modernes, mais n'offre pas l'isolation complète des conteneurs actuels.

## Cgroups (Control Groups)

Les cgroups sont un mécanisme qui permet de limiter, comptabiliser et isoler l'utilisation des ressources système (CPU, mémoire, E/S disque, réseau, etc.) pour un ensemble de processus.

### Fonctionnalités principales :
- **Limitation des ressources** : définir des limites de CPU, mémoire, etc.
- **Priorisation** : allouer plus ou moins de ressources à certains processus
- **Comptabilisation** : mesurer la consommation de ressources
- **Contrôle** : isoler et gérer des groupes de processus ensemble

Les cgroups permettent de s'assurer qu'un conteneur ne peut pas monopoliser toutes les ressources du système hôte.

## Fonctionnement combiné

Dans un conteneur moderne comme Docker :
- Les **namespaces** fournissent l'isolation au niveau du système d'exploitation
- **Chroot** (ou des mécanismes similaires plus avancés) assure l'isolation du système de fichiers
- Les **cgroups** garantissent que le conteneur utilise uniquement les ressources qui lui sont allouées

Cette combinaison permet d'avoir des conteneurs légers qui partagent le même noyau Linux mais fonctionnent comme des environnements isolés.