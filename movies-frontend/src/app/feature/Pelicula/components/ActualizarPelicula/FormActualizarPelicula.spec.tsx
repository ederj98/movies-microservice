import * as React from 'react';
import { RenderResult, fireEvent, render, wait } from '@testing-library/react';
import { SinonStub, stub } from 'sinon';
import { FormActualizarPelicula } from './';
import { setTextEvent } from '../../../../shared/utils/test';

describe('FormActualizarPelicula test', () => {
  let componentWrapper: RenderResult;
  let componentProps: React.ComponentProps<typeof FormActualizarPelicula> & {
    onSubmit: SinonStub;
  };

  beforeEach(() => {
    componentProps = {
      pelicula: {id: 1, name: 'Interstellar', director: 'The Grinch', writer: 'John Doe', stars: 'John Doe, Jane Doe'},
      formTitle: 'Form test',
      onSubmit: stub(),
    };
    componentWrapper = render(<FormActualizarPelicula {...componentProps} />);
  });

  it('should match snapshot', () => {
    expect(componentWrapper.container).toMatchSnapshot();
  });

  it('should submit with name updated', async () => {
    const elem = componentWrapper.container;
    const name = elem.querySelector('input[name="name"]');
    const submitButton = elem.querySelector('button[type="submit"]');

    await wait(() => {
      name && fireEvent.change(name, setTextEvent('name', 'Stargate'));
    });
    await wait(() => {
      submitButton && fireEvent.click(submitButton);
    });

    const formSubmitted = componentProps.onSubmit.firstCall.args[0];
    
    expect(formSubmitted.name).toBe('Stargate');
    expect(formSubmitted.director).toBe('The Grinch');
    expect(formSubmitted.writer).toBe('John Doe');
    expect(formSubmitted.stars).toBe('John Doe, Jane Doe');
  });

  it('should submit with writer updated', async () => {
    const elem = componentWrapper.container;
    const writer = elem.querySelector('input[name="writer"]');
    const submitButton = elem.querySelector('button[type="submit"]');

    await wait(() => {
      writer && fireEvent.change(writer, setTextEvent('writer', 'Maria Jane Watson'));
    });
    await wait(() => {
      submitButton && fireEvent.click(submitButton);
    });

    const formSubmitted = componentProps.onSubmit.firstCall.args[0];
    
    expect(formSubmitted.name).toBe('Interstellar');
    expect(formSubmitted.director).toBe('The Grinch');
    expect(formSubmitted.writer).toBe('Maria Jane Watson');
    expect(formSubmitted.stars).toBe('John Doe, Jane Doe');
  });

  it('should submit with stars updated', async () => {
    const elem = componentWrapper.container;
    const stars = elem.querySelector('input[name="stars"]');
    const submitButton = elem.querySelector('button[type="submit"]');

    await wait(() => {
      stars && fireEvent.change(stars, setTextEvent('stars', 'Miguel J Jhonson, Caprica Six'));
    });
    await wait(() => {
      submitButton && fireEvent.click(submitButton);
    });

    const formSubmitted = componentProps.onSubmit.firstCall.args[0];
    
    expect(formSubmitted.name).toBe('Interstellar');
    expect(formSubmitted.director).toBe('The Grinch');
    expect(formSubmitted.writer).toBe('John Doe');
    expect(formSubmitted.stars).toBe('Miguel J Jhonson, Caprica Six');
  });

  it('should submit director updated', async () => {
    const elem = componentWrapper.container;
    const director = elem.querySelector('input[name="director"]');
    const submitButton = elem.querySelector('button[type="submit"]');

    await wait(() => {
      director && fireEvent.change(director, setTextEvent('director', 'Jane M Doe'));
    });
    await wait(() => {
      submitButton && fireEvent.click(submitButton);
    });

    const formSubmitted = componentProps.onSubmit.firstCall.args[0];
    
    expect(formSubmitted.name).toBe('Interstellar');
    expect(formSubmitted.director).toBe('Jane M Doe');
    expect(formSubmitted.writer).toBe('John Doe');
    expect(formSubmitted.stars).toBe('John Doe, Jane Doe');
  });
});
