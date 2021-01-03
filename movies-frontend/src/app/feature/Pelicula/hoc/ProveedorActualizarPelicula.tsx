import {
  buscarPeliculaAsync,
  actualizarPeliculaAsync, 
} from '../../../core/redux/acciones/peliculas/PeliculasAcciones';
import { EstadoGeneral } from '../../../core/redux/modelo/EstadoGeneral';
import { ActualizarPelicula } from '../containers/GestionPeliculas/Actualizar';
import { connect } from 'react-redux';

const mapStateToProps = (state: EstadoGeneral) => {
  return  state.peliculas ;
};

export const ProveedorActualizarPelicula = connect(mapStateToProps, {
  buscarPelicula: buscarPeliculaAsync,
  actualizarPelicula: actualizarPeliculaAsync,
})(ActualizarPelicula);
