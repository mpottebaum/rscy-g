# make dir if no existo
mkdir -p dist

# copy deps metadata
cp package.json dist/package.json
cp package-lock.json dist/package-lock.json

# get in that dir, bb
cd dist

# reset type to default
npm pkg set 'type'='commonjs'

# captainjack install depps production style
npm ci --production
