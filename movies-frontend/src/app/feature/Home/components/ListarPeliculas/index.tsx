import * as PropTypes from 'prop-types';
import * as React from 'react';
import { BtnEliminarPelicula } from '../EliminarPelicula';
import { BtnActualizarPelicula } from '../ActualizarPelicula'
import { Pelicula } from '../../../Pelicula/models/Pelicula';
import { Table } from './styles';

export interface ListaPeliculasProps {
  peliculas: Array<Pelicula>;
  onClickEliminarPelicula: (peliculas: Pelicula) => void;
}

export const ListaPeliculas: React.FC<ListaPeliculasProps> = ({
  peliculas,
  onClickEliminarPelicula,
}) => {
  return (
    <Table>
      <thead>
        <tr>
          <td>
            <b>Name</b>
          </td>
          <td>
            <b>Director</b>
          </td>
          <td>
            <b>Writer</b>
          </td>
          <td>
            <b>Stars</b>
          </td>
          <td>
            <b>Eliminar</b>
          </td>
          <td>
            <b>Actualizar</b>
          </td>
        </tr>
      </thead>
      <tbody>
        {peliculas.map((pelicula: Pelicula) => {
          return (
            <tr key={pelicula.id}>
              <td>{pelicula.name}</td>
              <td>{pelicula.director}</td>
              <td>{pelicula.writer}</td>
              <td>{pelicula.stars}</td>
              <td>
                <BtnEliminarPelicula
                  pelicula={pelicula}
                  onEliminar={onClickEliminarPelicula}
                ></BtnEliminarPelicula>
              </td>
              <td>
                <BtnActualizarPelicula
                  pelicula={pelicula}
                  onActualizar={onClickEliminarPelicula}
                ></BtnActualizarPelicula>
              </td>
            </tr>
          );
        })}
      </tbody>
    </Table>
  );
};

ListaPeliculas.propTypes = {
  peliculas: PropTypes.array.isRequired,
  onClickEliminarPelicula: PropTypes.func.isRequired,
};
