import React, { useState } from 'react';
import './sidenav.css';
import { Link, useLocation } from 'react-router-dom';
import InventoryIcon from '@mui/icons-material/Inventory';
import ReceiptIcon from '@mui/icons-material/Receipt';
import AssessmentIcon from '@mui/icons-material/Assessment';
import BadgeIcon from '@mui/icons-material/Badge';
import MenuIcon from '@mui/icons-material/Menu';
import MenuOpenIcon from '@mui/icons-material/MenuOpen';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import LogoutIcon from '@mui/icons-material/Logout';

const sideNavMenu = [
  {
    icon: InventoryIcon,
    name: 'inventory',
    path: '/inventory'
  },
  {
    icon: ReceiptIcon,
    name: 'transaction',
    path: '/transaction'
  },
  {
    icon: AssessmentIcon,
    name: 'report',
    path: '/report'
  },
  {
    icon: BadgeIcon,
    name: 'employee',
    path: '/employee'
  }
];

function SideNav() {
  const location = useLocation();

  const [isWide, setIsWide] = useState<boolean>(window.innerWidth > 1024);
  const [isActive, setIsActive] = useState<string>(location.pathname);

  function handleClick(path: string) {
    setIsActive(path);
  }

  return (
    <div className={`sidenav ${isWide ? '' : 'short'}`}>
      <div
        className={`sidenav-hamburg ${isWide ? '' : 'short'}`}
        onClick={() => { setIsWide(!isWide); }}
        onKeyDown={() => {}}
        role="button"
        tabIndex={0}
      >
        {
          isWide ? <MenuOpenIcon sx={{ fontSize: 40 }} /> : <MenuIcon sx={{ fontSize: 40 }} />
        }
      </div>
      <div className={`sidenav-header ${isWide ? '' : 'short'}`}>
        <div>g.i.m.s.</div>
      </div>
      <div className="sidenav-menu">
        {
          sideNavMenu.map((data) => (
            <Link
              className={`sidenav-menu-item ${isActive === data.path ? 'active' : ''}`}
              to={data.path}
              onClick={() => { handleClick(data.path); }}
              key={data.name}
            >
              <data.icon className="icon" />
              <div className="name">{data.name}</div>
              <div className="tooltip">{data.name}</div>
            </Link>
          ))
        }
      </div>
      <div className="sidenav-account">
        <div className="account-box">
          <AccountCircleIcon className="account-icon" />
          <div className="account-name">Nik Kunraho Struyf</div>
        </div>
        <div className="account-logout">
          <LogoutIcon className="logout-icon" />
          <div className="tooltip">logout</div>
        </div>
      </div>
    </div>
  );
}

export default SideNav;