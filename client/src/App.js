import React, { Fragment } from 'react';

import './App.css';
import Header from './components/Header';
import Footer from './components/Footer';
import Routes from './Routes';

function App() {
  return (
    <Fragment>
      <Header />
        <main className="container">
          <Routes />
        </main>
      <Footer />
    </Fragment>
  );
}

export default App;
