// const ApiConfig = {
//   tmdb_ApiKey: 'c6e09799f073fdcab59c60b86ee5d712',
//   tmdb_baseUrl: 'https://api.themoviedb.org/3',
//   tmdb_originalImage: (imgPath) => `https://image.tmdb.org/t/p/original/${imgPath}`,
//   tmdb_w500Image: (imgPath) => `https://image.tmdb.org/t/p/w500/${imgPath}`,
// };

// export default ApiConfig;

const config = {
  path_file: (img) => `http://localhost:5000/uploads/${img}`,
};

export default config;
