//const path = require('path');
const express = require('express');
const { exec } = require('child_process');
const app = express();
const port = 3000;

//const golangFilePath = 'C:/Users/Harshal Trivedi/cloud-1/my-node-app/cli.go';

app.get('/api', (req, res) => {
 exec(`go run cli.go`, (error, stdout, stderr) => { //${golangFilePath}
    if (error) {
        console.log(`error: ${error.message}`);
        return;
      }
      if (stderr) {
        console.log(`stderr: ${stderr}`);
        return;
      }
      console.log(`stdout: ${stdout}`);
      res.json({ message: 'Golang code executed successfully' });
    });
  });

  app.listen(port, () => {
    console.log(`Server listening on port ${port}`);
 });