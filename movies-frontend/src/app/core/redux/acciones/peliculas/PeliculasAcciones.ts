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
    ).catch(error => {
      alert('Ocurrio un error al consultar las Peliculas')
      console.error('Error!', error.response);
    });
  };
}

export function buscarPeliculaAsync(id: number) {
  return function (dispacth: any) {
    PeliculaRepositorio.consultar(id)
    .then((respuesta: any) =>
      dispacth(
        buscarPelicula(respuesta.data)
      )
    ).catch(error => {
      if (error.response.status === 400) {
        alert('Error en formato del ID, debe ser numerico')
      } else if (error.response.status === 404) {
          alert('Pelicula no registrada!')
      } else {
        alert('Ocurrido un error al consultar la Pelicula')
      }
      console.error('Error!', error.response);
    });
  };
}

export function agregarNuevaPeliculaAsync(pelicula: Pelicula) {
  return function (dispacth: any) {
    PeliculaRepositorio.guardar(pelicula)
    .then((respuesta: any) => {
      alert('La Pelicula se guardo correctamente!')
      dispacth(
        agregarNuevaPelicula(pelicula)
      )
    }
    ).catch(error => {
      if (error.response.status === 400 && 
            error.response.data.message === 'The movie is already exist') {
        alert('Error! La Pelicula ya se encuentra registrada!')
      } else {
        alert('Error! Ocurrido un error al crear la Pelicula')
      }
      console.error('Error!', error.response.data);
    });
  };
}

export function actualizarPeliculaAsync(pelicula: Pelicula) {
  return function (dispacth: any) {
    PeliculaRepositorio.actualizar(pelicula.id, pelicula)
    .then((respuesta: any) =>
      dispacth(
        actualizarPelicula(pelicula)
      )
    ).catch(error => {
      console.error('Error!', error.response.data);
    });
  };
}

export function eliminarPeliculaAsync(pelicula: Pelicula) {
  return function (dispacth: any) {
    PeliculaRepositorio.eliminar(pelicula.id)
    .then((respuesta: any) => {
      alert('La Pelicula fue eliminada correctamente')
      dispacth(
        eliminarPelicula(pelicula)
      )
    }
    ).catch(error => {
      if (error.response.status === 404) {
          alert('Pelicula no registrada!')
      } else {
        alert('Ocurrido un error al eliminar la Pelicula')
      }
      console.error('Error!', error.response);
    });
  };
}
