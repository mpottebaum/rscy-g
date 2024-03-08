import nodeResolve from '@rollup/plugin-node-resolve';
import commonjs from '@rollup/plugin-commonjs';
import babel from '@rollup/plugin-babel';
import replace from '@rollup/plugin-replace';
import typescript from '@rollup/plugin-typescript';

export default {
   input: 'src/gates-of-rsc.tsx',
   output: {
      file: 'dist/gates-of-rsc.js',
      format: 'iife',
   },
   plugins: [
      nodeResolve(
        {
         extensions: ['.ts', '.tsx'],
         mainFields: ['module'],
      }
      ),
      commonjs(),
      babel({
         babelHelpers: 'bundled',
         presets: ['@babel/preset-react',],
         extensions: ['.ts', '.tsx']
      }),
      typescript({ tsconfig: './tsconfig.json' }),
      replace({
         preventAssignment: false,
         'process.env.NODE_ENV': '"production"'
      }),
   ]
}