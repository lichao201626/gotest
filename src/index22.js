import React from 'react';
import { render } from 'react-dom';
import App from './App'

const renderDom = Component => {
    render(
        <App/>,
        document.getElementById('example')
    );
}
renderDom(App);