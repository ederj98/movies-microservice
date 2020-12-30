import { Pelicula } from 'app/feature/Pelicula/models/Pelicula';
import { axiosIntance } from '../config/AxiosConfig';

export const PeliculaRepositorio = {
  consultarTodos: () =>
    axiosIntance.get(`/api/movies`),
  
  consultar: (id: number) => 
    axiosIntance.get(`/api/movies/${id}`),
  
  guardar: (pelicula: Pelicula) =>
    axiosIntance.post(`/api/movies`, pelicula),

  actualizar: (id: number, pelicula: Pelicula) =>
    axiosIntance.put(`/api/movies/${id}`, pelicula),

  eliminar: (id: number) => 
    axiosIntance.delete(`/api/movies/${id}`)
};
