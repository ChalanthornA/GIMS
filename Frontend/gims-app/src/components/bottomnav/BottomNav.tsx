import React, { useState } from 'react';
import './bottomnav.css';
import { Link, useLocation } from 'react-router-dom';
import InventoryIcon from '@mui/icons-material/Inventory';
import ReceiptIcon from '@mui/icons-material/Receipt';
import AssessmentIcon from '@mui/icons-material/Assessment';
import BadgeIcon from '@mui/icons-material/Badge';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';

const BottomNavMenu = [
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
  },
  {
    icon: AccountCircleIcon,
    name: 'account',
    path: '/account'
  }
];

function BottomNav() {
  const location = useLocation();

  const [isActive, setIsActive] = useState<string>(location.pathname);

  function handleClick(path: string) {
    setIsActive(path);
  }

  return (
    <div className="bottomnav">
      <div className="bottomnav-menu">
        {
          BottomNavMenu.map((data) => (
            <Link
              className={`bottomnav-menu-item ${isActive === data.path ? 'active' : ''}`}
              to={data.path}
              onClick={() => { handleClick(data.path); }}
              key={data.name}
            >
              <data.icon className="bottomnav-item-icon" sx={{ fontSize: '3.5vh' }} />
              <div className="bottomnav-item-name">{data.name}</div>
            </Link>
          ))
        }
      </div>
    </div>
  );
}

export default BottomNav;
