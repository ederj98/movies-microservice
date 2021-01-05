import * as React from 'react';
import { RenderResult, fireEvent, render, wait } from '@testing-library/react';
import { SinonStub, stub } from 'sinon';
import { FormCrearPelicula } from '.';
import { setTextEvent } from '../../../../shared/utils/test';

describe('FormCrearPelicula test', () => {
  let componentWrapper: RenderResult;
  let componentProps: React.ComponentProps<typeof FormCrearPelicula> & {
    onSubmit: SinonStub;
  };

  beforeEach(() => {
    componentProps = {
      formTitle: 'Form test',
      onSubmit: stub(),
    };
    componentWrapper = render(<FormCrearPelicula {...componentProps} />);
  });

  it('should match snapshot', () => {
    expect(componentWrapper.container).toMatchSnapshot();
  });

  it('should fail on submit all fields missing', async () => {
    const elem = componentWrapper.container;
    const submitButton = elem.querySelector('button[type="submit"]');

    await wait(() => {
      submitButton && fireEvent.click(submitButton);
    });
    const spans = elem.querySelectorAll('span');
    expect(spans.length).toBe(4);
    expect(spans[0].textContent).toBe('El campo Name es requerido.');
    expect(spans[1].textContent).toBe('El campo Director es requerido.');
    expect(spans[2].textContent).toBe('El campo Writer es requerido.');
    expect(spans[3].textContent).toBe('El campo Stars es requerido.');
  });

  it('should fail on submit three fields missing', async () => {
    const elem = componentWrapper.container;
    const name = elem.querySelector('input[name="name"]');
    const submitButton = elem.querySelector('button[type="submit"]');

    await wait(() => {
      name && fireEvent.change(name, setTextEvent('name', 'Interstellar'));
    });

    await wait(() => {
      submitButton && fireEvent.click(submitButton);
    });
    const spans = elem.querySelectorAll('span');
    expect(spans.length).toBe(3);
    expect(spans[0].textContent).toBe('El campo Director es requerido.');
    expect(spans[1].textContent).toBe('El campo Writer es requerido.');
    expect(spans[2].textContent).toBe('El campo Stars es requerido.');
  });

  it('should fail on submit two fields missing', async () => {
    const elem = componentWrapper.container;
    const name = elem.querySelector('input[name="name"]');
    const director = elem.querySelector('input[name="director"]');
    const submitButton = elem.querySelector('button[type="submit"]');

    await wait(() => {
      name && fireEvent.change(name, setTextEvent('name', 'Interstellar'));
    });
    await wait(() => {
      director && fireEvent.change(director, setTextEvent('director', 'The Grinch'));
    });

    await wait(() => {
      submitButton && fireEvent.click(submitButton);
    });
    const spans = elem.querySelectorAll('span');
    expect(spans.length).toBe(2);
    expect(spans[0].textContent).toBe('El campo Writer es requerido.');
    expect(spans[1].textContent).toBe('El campo Stars es requerido.');
  });

  it('should fail on submit one field missing', async () => {
    const elem = componentWrapper.container;
    const name = elem.querySelector('input[name="name"]');
    const director = elem.querySelector('input[name="director"]');
    const writer = elem.querySelector('input[name="writer"]');
    const submitButton = elem.querySelector('button[type="submit"]');

    await wait(() => {
      name && fireEvent.change(name, setTextEvent('name', 'Interstellar'));
    });
    await wait(() => {
      director && fireEvent.change(director, setTextEvent('director', 'The Grinch'));
    });
    await wait(() => {
      writer && fireEvent.change(writer, setTextEvent('writer', 'John Doe'));
    });

    await wait(() => {
      submitButton && fireEvent.click(submitButton);
    });
    const spans = elem.querySelectorAll('span');
    expect(spans.length).toBe(1);
    expect(spans[0].textContent).toBe('El campo Stars es requerido.');
  });

  it('should submit', async () => {
    const elem = componentWrapper.container;

    const name = elem.querySelector('input[name="name"]');
    const director = elem.querySelector('input[name="director"]');
    const writer = elem.querySelector('input[name="writer"]');
    const stars = elem.querySelector('input[name="stars"]');
    const submitButton = elem.querySelector('button[type="submit"]');

    await wait(() => {
      name && fireEvent.change(name, setTextEvent('name', 'Interstellar'));
    });
    await wait(() => {
      director && fireEvent.change(director, setTextEvent('director', 'The Grinch'));
    });
    await wait(() => {
      writer && fireEvent.change(writer, setTextEvent('writer', 'John Doe'));
    });
    await wait(() => {
      stars && fireEvent.change(stars, setTextEvent('stars', 'John Doe, Jane Doe'));
    });

    await wait(() => {
      submitButton && fireEvent.click(submitButton);
    });

    const formSubmitted = componentProps.onSubmit.firstCall.args[0];
    
    expect(formSubmitted.name).toBe('Interstellar');
    expect(formSubmitted.director).toBe('The Grinch');
    expect(formSubmitted.writer).toBe('John Doe');
    expect(formSubmitted.stars).toBe('John Doe, Jane Doe');
  });
});
