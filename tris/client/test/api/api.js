(async () => {
  let curr =  Date.now();
  const rawResponse = await fetch('localhost:8080', {
    method: 'POST',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({time: curr, })
  });
  const content = await rawResponse.json();

  console.log(content);
})();