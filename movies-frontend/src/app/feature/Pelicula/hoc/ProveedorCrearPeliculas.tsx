import {
  agregarNuevaPeliculaAsync,
} from '../../../core/redux/acciones/peliculas/PeliculasAcciones';
import { EstadoGeneral } from '../../../core/redux/modelo/EstadoGeneral';
import { GestionPeliculas } from '../containers/GestionPeliculas/Crear';
import { connect } from 'react-redux';

const mapStateToProps = (state: EstadoGeneral) => {
  return state.peliculas;
};

export const ProveedorCrearPeliculas = connect(mapStateToProps, {
  agregarNuevaPelicula: agregarNuevaPeliculaAsync,
})(GestionPeliculas);
