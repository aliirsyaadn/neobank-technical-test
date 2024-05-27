// src/components/Layout.js
import React from 'react';
import { Outlet } from 'react-router-dom';
import SideMenu from './SideMenu';
import TopMenu from './TopMenu';

const Layout = () => {
  return (
    <div className="flex">
      <SideMenu />
      <div
        className="content bg-gray-300"
        style={{
          width: 'calc(100vw - 208px)'
        }}
      >
        <TopMenu />
        <Outlet/>
      </div>
    </div>
  );
};

export default Layout;
