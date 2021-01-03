import {
  AGREGAR_PELICULA,
  ACTUALIZAR_PELICULA,
  BUSCAR_PELICULA,
  ELIMINAR_PELICULA,
  LISTAR_PELICULAS,
  TiposAccionesPelicula,
} from './PeliculasTiposAcciones';
import { Pelicula } from 'app/feature/Pelicula/models/Pelicula';
import { PeliculaRepositorio } from 'app/core/api/peliculas.repositorio';

export function listarPeliculas(
  peliculas: Array<Pelicula>
): TiposAccionesPelicula {
  return {
    type: LISTAR_PELICULAS,
    payload: peliculas,
  };
}

export function agregarNuevaPelicula(
  pelicula: Pelicula
): TiposAccionesPelicula {
  return {
    type: AGREGAR_PELICULA,
    payload: pelicula,
  };
}

export function buscarPelicula(
  pelicula: Pelicula
): TiposAccionesPelicula {
  console.log('buscar')
  return {
    type: BUSCAR_PELICULA,
    payload: pelicula,
  };
}

export function actualizarPelicula(
  pelicula: Pelicula
): TiposAccionesPelicula {
  return {
    type: ACTUALIZAR_PELICULA,
    payload: pelicula,
  };
}

export function eliminarPelicula(pelicula: Pelicula): TiposAccionesPelicula {
  return {
    type: ELIMINAR_PELICULA,
    payload: pelicula,
  };
}

export function listarPeliculasAsync() {
  return function (dispacth: any) {
    PeliculaRepositorio.consultarTodos()
    .then((respuesta: any) =>
      dispacth(
        listarPeliculas(respuesta.data)
      )
    );
  };
}

export function buscarPeliculaAsync(id: number) {
  return function (dispacth: any) {
    PeliculaRepositorio.consultar(id)
    .then((respuesta: any) =>
      dispacth(
        buscarPelicula(respuesta.data)
      )
    );
  };
}

export function agregarNuevaPeliculaAsync(pelicula: Pelicula) {
  return function (dispacth: any) {
    PeliculaRepositorio.guardar(pelicula)
    .then((respuesta: any) =>
      dispacth(
        agregarNuevaPelicula(pelicula)
      )
    ).catch(error => {
      console.error('There was an error!', error);
    });
  };
}

export function actualizarPeliculaAsync(pelicula: Pelicula) {
  return function (dispacth: any) {
    PeliculaRepositorio.actualizar(pelicula.Id, pelicula)
    .then((respuesta: any) =>
      dispacth(
        actualizarPelicula(pelicula)
      )
    ).catch(error => {
      console.error('There was an error!', error);
    });
  };
}

export function eliminarPeliculaAsync(pelicula: Pelicula) {
  return function (dispacth: any) {
    PeliculaRepositorio.eliminar(pelicula.Id)
    .then((respuesta: any) =>
      dispacth(
        eliminarPelicula(pelicula)
      )
    ).catch(error => {
      console.error('There was an error!', error);
    });
  };
}
