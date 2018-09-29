import React from 'react';
import { render } from 'react-dom';
import App from './App'

const renderDom = Component => {
    render(
        <h1>Hello, world!</h1>,
        document.getElementById('example')
    );
}
renderDom(App);