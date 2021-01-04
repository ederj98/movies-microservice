import {
  eliminarPeliculaAsync,
  listarPeliculasAsync,
} from '../../../core/redux/acciones/peliculas/PeliculasAcciones';
import { EstadoGeneral } from '../../../core/redux/modelo/EstadoGeneral';
import { GestionListarPeliculas } from '../containers/ListarPeliculas';
import { connect } from 'react-redux';

const mapStateToProps = (state: EstadoGeneral) => {
  return state.peliculas;
};

export const ProveedorListarGestionPeliculas = connect(mapStateToProps, {
  listarPeliculas: listarPeliculasAsync,
  eliminarPelicula: eliminarPeliculaAsync,
})(GestionListarPeliculas);
