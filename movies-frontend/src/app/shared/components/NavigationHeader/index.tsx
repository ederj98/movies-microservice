import * as React from 'react';
import { HeaderNav } from './styles';
import Logo from 'assets/img/movies-logo.png';
import { NavBrand } from './NavBrand';
import { NavList } from './NavList';

export const NavigationHeader: React.FC = () => {
  const routes = [
    { label: 'Home', url: '/home' },
    { label: 'Movies', url: '/movie' },
  ];
  return (
    <HeaderNav>
      <NavBrand imgSrc={Logo} text="Movies"></NavBrand>
      <NavList items={routes} />
    </HeaderNav>
  );
};
