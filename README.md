# Comprendre les Containers "From Scratch"

Ce document explique les concepts fondamentaux et la démarche pour construire un container depuis zéro, basé sur la présentation de Liz Rice.

## Table des matières
- [Introduction](#introduction)
- [Concepts clés](#concepts-clés)
  - [Namespaces](#namespaces)
  - [Chroot](#chroot)
  - [Control Groups (cgroups)](#control-groups-cgroups)
- [Démarche pour construire un container](#démarche-pour-construire-un-container)
  - [Point de départ](#point-de-départ)
  - [Implémentation des Namespaces](#implémentation-des-namespaces)
  - [Isolation du système de fichiers](#isolation-du-système-de-fichiers)
  - [Limitation des ressources avec cgroups](#limitation-des-ressources-avec-cgroups)
## Introduction

Un container est une unité logicielle qui encapsule une application et ses dépendances, permettant son exécution de manière isolée et portable. La présentation de Liz Rice démontre comment construire un container en utilisant directement les fonctionnalités du noyau Linux, sans s'appuyer sur Docker ou d'autres outils de containerisation.

## Concepts clés

### Namespaces

Les namespaces fournissent une **vue isolée** des ressources du système pour un groupe de processus.

Types de namespaces principaux :
- **UTS (Unix Time-sharing System)** : Isole le nom d'hôte et le nom de domaine
- **PID (Process ID)** : Isole les identifiants de processus
- **Mount** : Isole les points de montage du système de fichiers
- **Network** : Isole les interfaces réseau, les adresses IP, les ports
- **User** : Isole les identifiants d'utilisateurs et de groupes
- **IPC (Inter-Process Communication)** : Isole les mécanismes de communication entre processus

### Chroot

`chroot` est une opération qui **change le répertoire racine apparent** pour un processus et ses enfants. Cela permet d'isoler le système de fichiers d'un container du reste du système.

### Control Groups (cgroups)

Les cgroups permettent de **limiter, comptabiliser et isoler l'utilisation des ressources** (CPU, mémoire, E/S, nombre de processus) pour des groupes de processus spécifiques.

Fonctionnalités principales :
- Limitation des ressources (mémoire, CPU, etc.)
- Priorisation des ressources
- Comptabilisation de l'utilisation
- Contrôle des processus

## Démarche pour construire un container

### Point de départ

La démarche commence par un programme simple en Go qui exécute une commande arbitraire, similaire à `docker run`.

### Implémentation des Namespaces

1. **UTS Namespace (Nom d'hôte)**
   - Utilisation de `syscall.CLONE_NEWUTS` pour créer un nouveau namespace UTS
   - Mise en place d'un processus parent ("runner") et d'un processus enfant ("child")
   - Configuration du nom d'hôte personnalisé dans le container

2. **PID Namespace (Identifiants de Processus)**
   - Ajout de `syscall.CLONE_NEWPID` pour isoler les PIDs
   - Les processus à l'intérieur du container commencent à l'ID 1
   - Configuration des systèmes de fichiers pseudo comme `/proc` pour que les outils comme `ps` fonctionnent correctement

3. **Mount Namespace (Isolation des Montages)**
   - Utilisation de `syscall.CLONE_NEWNS` pour isoler les points de montage
   - Assurance que les montages du container ne soient pas visibles sur l'hôte

### Isolation du système de fichiers

1. Utilisation de `syscall.Chroot` pour changer la racine du processus containerisé
2. Préparation d'un répertoire avec le système de fichiers nécessaire (`vagrant-ubuntu-fs`)
3. Montage du système de fichiers pseudo `/proc` à l'intérieur du container pour le bon fonctionnement des outils

### Limitation des ressources avec cgroups

1. Création de répertoires dans `/sys/fs/cgroup/` pour le container
2. Configuration des limites de ressources en écrivant dans des fichiers spécifiques :
   - `memory.limit_in_bytes` pour limiter la mémoire
   - `pids.max` pour limiter le nombre de processus
3. Ajout du PID du processus containerisé à `cgroup.procs` pour appliquer ces limites