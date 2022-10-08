import React, { useState } from 'react';
import './sidenav.css';
import InventoryIcon from '@mui/icons-material/Inventory';
import ReceiptIcon from '@mui/icons-material/Receipt';
import AssessmentIcon from '@mui/icons-material/Assessment';
import BadgeIcon from '@mui/icons-material/Badge';
import MenuIcon from '@mui/icons-material/Menu';
import MenuOpenIcon from '@mui/icons-material/MenuOpen';

const sideNavMenu = [
  {
    icon: InventoryIcon,
    name: 'inventory'
  },
  {
    icon: ReceiptIcon,
    name: 'transaction'
  },
  {
    icon: AssessmentIcon,
    name: 'report'
  },
  {
    icon: BadgeIcon,
    name: 'employee'
  }
];

function SideNav() {
  const [isWide, setIsWide] = useState<boolean>(true);
  const [isActive, setIsActive] = useState<string>('inventory');

  function handleClick(name: string) {
    setIsActive(name);
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
      <div className="sidenav-header">g.i.m.s.</div>
      <div className="sidenav-menu">
        {
          sideNavMenu.map((data) => (
            <a
              className={`sidenav-menu-item ${isActive === data.name ? 'active' : ''}`}
              href="/"
              onClick={() => { handleClick(data.name); }}
            >
              <data.icon className="icon" />
              <div className="name">{data.name}</div>
            </a>
          ))
        }
      </div>
    </div>
  );
}

export default SideNav;
