import {
  agregarNuevaPeliculaAsync,
  eliminarPelicula,
  listarPeliculasAsync,
} from '../../../core/redux/acciones/peliculas/PeliculasAcciones';
import { EstadoGeneral } from '../../../core/redux/modelo/EstadoGeneral';
import { GestionPeliculas } from '../containers/GestionPeliculas';
import { connect } from 'react-redux';

const mapStateToProps = (state: EstadoGeneral) => {
  return state.peliculas;
};

export const ProveedorGestionPeliculas = connect(mapStateToProps, {
  listarPeliculas: listarPeliculasAsync,
  agregarNuevaPelicula: agregarNuevaPeliculaAsync,
  eliminarPelicula,
})(GestionPeliculas);
