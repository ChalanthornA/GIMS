import React from 'react';
import './signinpage.css';

function SignInPage() {
  return (
    <div className="background">
      <div className="container-signin">
        <div className="header-signin">sign in</div>
        <input className="input-userid" type="text" placeholder="USER ID" />
        <input className="input-passwd" type="password" placeholder="PASSWORD" />
        <a className="link-forgotpass" href="/">forgot password?</a>
        <div className="input-incorrect">invalid username or password</div>
        <button
          className="button-signin"
          type="submit"
          onClick={() => {}}
          onKeyDown={() => {}}
        >
          sign in
        </button>
      </div>
    </div>
  );
}

export default SignInPage;
