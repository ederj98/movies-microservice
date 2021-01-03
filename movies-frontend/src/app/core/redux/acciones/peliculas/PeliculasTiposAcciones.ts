import { Pelicula } from 'app/feature/Pelicula/models/Pelicula';

export const LISTAR_PELICULAS = 'LISTAR_PELICULAS';
export const AGREGAR_PELICULA = 'AGREGAR_PELICULA';
export const BUSCAR_PELICULA = 'BUSCAR_PELICULA';
export const ACTUALIZAR_PELICULA = 'ACTUALIZAR_PELICULA';
export const ELIMINAR_PELICULA = 'ELIMINAR_PELICULA';

interface AccionListarPeliculas {
  type: typeof LISTAR_PELICULAS;
  payload: Pelicula[];
}

interface AccionAgregarPelicula {
  type: typeof AGREGAR_PELICULA;
  payload: Pelicula;
}

interface AccionBuscarPelicula {
  type: typeof BUSCAR_PELICULA;
  payload: Pelicula;
}

interface AccionActualizarPelicula {
  type: typeof ACTUALIZAR_PELICULA;
  payload: Pelicula;
}

interface AccionEliminarPelicula {
  type: typeof ELIMINAR_PELICULA;
  payload: Pelicula;
}

export type TiposAccionesPelicula =
  | AccionListarPeliculas
  | AccionAgregarPelicula
  | AccionBuscarPelicula
  | AccionActualizarPelicula
  | AccionEliminarPelicula;
