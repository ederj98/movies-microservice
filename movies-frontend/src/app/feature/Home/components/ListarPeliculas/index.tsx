import * as PropTypes from 'prop-types';
import * as React from 'react';
import { BtnEliminarPelicula } from '../EliminarPelicula';
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
        </tr>
      </thead>
      <tbody>
        {peliculas.map((pelicula: Pelicula) => {
          return (
            <tr key={pelicula.Id}>
              <td>{pelicula.Name}</td>
              <td>{pelicula.Director}</td>
              <td>{pelicula.Writer}</td>
              <td>{pelicula.Stars}</td>
              <td>
                <BtnEliminarPelicula
                  pelicula={pelicula}
                  onEliminar={onClickEliminarPelicula}
                ></BtnEliminarPelicula>
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
