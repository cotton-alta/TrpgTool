export default (roll) => {
  let countRe = /\d{1,2}/;
  let typeRe = /(d)(\d*)/;
  let count = countRe.exec(roll);
  let type = typeRe.exec(roll);
  let sum = 0;
  for(let i = 0; i < count[0]; i++) {
    sum += Math.floor(Math.random() * type[2]) + 1;
  }
  return sum;
};