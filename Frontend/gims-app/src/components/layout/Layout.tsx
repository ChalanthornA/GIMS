import React from 'react';
import './layout.css';
import { Outlet } from 'react-router-dom';
import SideNav from '../sidenav/SideNav';
import BottomNav from '../bottomnav/BottomNav';

function Layout() {
  return (
    <div className="layout">
      {
        window.innerWidth > 480 && <SideNav />
      }
      <div className="layout-content">
        <Outlet />
      </div>
      {
        window.innerWidth <= 480 && <BottomNav />
      }
    </div>
  );
}

export default Layout;
