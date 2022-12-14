ignite generate ts-client -y
cd ./ts-client
rm ./src/package.json
yarn install && yarn build
cp ./src/types.d.ts ./lib/types.d.ts
