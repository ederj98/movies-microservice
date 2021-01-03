import {
  AGREGAR_PELICULA,
  ELIMINAR_PELICULA,
  LISTAR_PELICULAS,
  BUSCAR_PELICULA,
  TiposAccionesPelicula,
} from '../../acciones/peliculas/PeliculasTiposAcciones';
import { EstadoPelicula } from '../../modelo/EstadoPelicula';
import { Pelicula } from 'app/feature/Pelicula/models/Pelicula';

const initialState: EstadoPelicula = {
  peliculas: Array<Pelicula>(),
  pelicula: {
    Id:0,
    Name: '',
    Director: '',
    Writer: '',
    Stars: '',
  }
};

export default function (
  state = initialState,
  action: TiposAccionesPelicula
): EstadoPelicula {
  switch (action.type) {
    case LISTAR_PELICULAS: {
      const peliculas = action.payload;
      return {
        ...state,
        peliculas,
      };
    }
    case BUSCAR_PELICULA: {
      const pelicula = action.payload;
      return {
        ...state,
        pelicula,
      };
    }
    case AGREGAR_PELICULA: {
      const pelicula = action.payload;
      return {
        ...state,
        peliculas: [...state.peliculas, pelicula],
      };
    }

    case ELIMINAR_PELICULA: {
      const pelicula = action.payload;
      return {
        ...state,
        peliculas: [
          ...state.peliculas.filter((p) => p.Name !== pelicula.Name),
        ],
      };
    }

    default:
      return state;
  }
}
