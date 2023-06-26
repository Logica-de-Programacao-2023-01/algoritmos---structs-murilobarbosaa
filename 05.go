package main

import "fmt"

type Musica struct {
	Titulo  string
	Artista string
	Duracao int
}

type Playlist struct {
	Nome    string
	Musicas []Musica
}

func buscarPlaylistsPorMusica(tituloMusica string, playlists []Playlist) []Playlist {
	var playlistsEncontradas []Playlist

	for _, playlist := range playlists {
		for _, musica := range playlist.Musicas {
			if musica.Titulo == tituloMusica {
				playlistsEncontradas = append(playlistsEncontradas, playlist)
				break
			}
		}
	}

	return playlistsEncontradas
}

func main() {
	playlist1 := Playlist{
		Nome: "Playlist 1",
		Musicas: []Musica{
			{Titulo: "Música 1", Artista: "Artista 1", Duracao: 180},
			{Titulo: "Música 2", Artista: "Artista 2", Duracao: 240},
		},
	}

	playlist2 := Playlist{
		Nome: "Playlist 2",
		Musicas: []Musica{
			{Titulo: "Música 2", Artista: "Artista 2", Duracao: 240},
			{Titulo: "Música 3", Artista: "Artista 3", Duracao: 210},
		},
	}

	playlist3 := Playlist{
		Nome: "Playlist 3",
		Musicas: []Musica{
			{Titulo: "Música 3", Artista: "Artista 3", Duracao: 210},
			{Titulo: "Música 4", Artista: "Artista 4", Duracao: 300},
		},
	}

	playlists := []Playlist{playlist1, playlist2, playlist3}

	titulo := "Música 2"
	playlistsEncontradas := buscarPlaylistsPorMusica(titulo, playlists)

	fmt.Printf("Playlists encontradas para a música %s:\n", titulo)
	for _, playlist := range playlistsEncontradas {
		fmt.Println(playlist.Nome)
	}
}
