import React from 'react';
import './layout.css';
import { Outlet } from 'react-router-dom';
import SideNav from '../sidenav/SideNav';
import BottomNav from '../bottomnav/BottomNav';

function Layout() {
  return (
    <div className="layout">
      {
        window.innerWidth > 500
          ? <SideNav /> : <BottomNav />
      }
      <div className="layout-content">
        <Outlet />
      </div>
    </div>
  );
}

export default Layout;
