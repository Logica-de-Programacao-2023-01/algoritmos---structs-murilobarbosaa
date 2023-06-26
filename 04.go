package main

import (
	"fmt"
	"time"
)

type Musica struct {
	Titulo  string
	Artista string
	Duracao time.Duration
}

type Playlist struct {
	Nome    string
	Musicas []Musica
}

func imprimirPlaylist(p Playlist) {
	fmt.Printf("Playlist: %s \n", p.Nome)

	totalDuracao := time.Duration(0)
	for _, musica := range p.Musicas {
		fmt.Printf("Título: %s \n", musica.Titulo)
		fmt.Printf("Artista: %s \n", musica.Artista)
		fmt.Printf("Duração: %s \n", musica.Duracao)

		totalDuracao += musica.Duracao
	}

	fmt.Printf("Tempo total da playlist: %s \n", totalDuracao)
}

func main() {
	musicas := []Musica{
		{Titulo: "Música 1", Artista: "Artista 1", Duracao: time.Minute*3 + time.Second*30},
		{Titulo: "Música 2", Artista: "Artista 2", Duracao: time.Minute*4 + time.Second*15},
		{Titulo: "Música 3", Artista: "Artista 3", Duracao: time.Minute*2 + time.Second*50},
	}

	playlist := Playlist{
		Nome:    "Minha playlist",
		Musicas: musicas,
	}

	imprimirPlaylist(playlist)
}
