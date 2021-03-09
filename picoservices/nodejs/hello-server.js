// ------------------------------------------------------------
// Giovanni Emanuele Nocco [Genocs].
// Licensed under the MIT License.
// ------------------------------------------------------------

const express = require('express');
const app = express();

const port = 3000;

app.get('/', (_req, res) => {
    console.log('Called home!!!');
    res.send('Hello from Node.js!!!');
});

app.get('/ping', (_req, res) => {
    res.send('pong');
});

app.listen(port, () => console.log(`Node App listening on port ${port}!`));