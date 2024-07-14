import fs from 'fs';
import { exec } from 'child_process';

// Read the user code
const code = fs.readFileSync('/usr/src/app/user_code.mjs', 'utf8');

// Create a new file that will import and render the user code
const scriptContent = `
import React from 'react';
import ReactDOMServer from 'react-dom/server';
${code}
const Component = React.createElement(Counter);
console.log(ReactDOMServer.renderToStaticMarkup(Component));
`;

// Write this script to a file
fs.writeFileSync('/usr/src/app/render.js', scriptContent);

// Run the script
exec('node --experimental-modules /usr/src/app/render.js', (error, stdout, stderr) => {
  if (error) {
    console.error('Error executing script:', stderr);
    process.exit(1);
  } else {
    console.log(stdout);
  }
});
