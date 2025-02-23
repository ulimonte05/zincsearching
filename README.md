# 🔍 ZincSearching

**Aplicación para buscar en datos de correos electrónicos usando ZincSearch, Go y Vue.js**

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Docker](https://badgen.net/badge/icon/docker?icon=docker&label)](https://www.docker.com)

![Demo](https://via.placeholder.com/800x400.png?text=ZincSearching+Interface+Preview) <!-- Agrega una imagen real aquí -->

Una solución moderna para indexar y buscar en grandes volúmenes de correos electrónicos, con:
- **Backend en Go** para procesamiento eficiente
- **Interfaz en Vue.js** intuitiva y responsiva
- **ZincSearch** como motor de búsqueda full-text

## 🚀 Comenzando

### Prerrequisitos
- [Docker](https://www.docker.com/get-started) instalado
- 4 GB de RAM disponibles
- Puertos 4080, 8080 y 5173 libres

### Instalación
```bash
1. Clona el repositorio:
git clone https://github.com/ulimonte05/zincsearching.git
cd zincsearching

2. Inicia los contenedores:
docker-compose up --build

```bash

### 🌍 Servicios desplegados
Servicio	Puerto	Descripción
ZincSearch	4080	Motor de búsqueda
API Go	8080	Backend REST
Client Vue.js	5173	Interfaz web

### 🧠 Datos de ejemplo incluidos
El sistema viene preconfigurado con:
10,000+ correos de ejemplo (dataset Enron 2011)

Índices pregenerados para búsquedas inmediatas

user: admin
pass:admin
