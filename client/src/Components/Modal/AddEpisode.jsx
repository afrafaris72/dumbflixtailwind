import { useContext, useState } from 'react';
import { useMutation, useQuery } from 'react-query';
import { API } from '../../Config/Api';
import { AlertSuccess, AlertError } from './AlertCollection';
import { ModalContext } from '../../Context/ModalContext';

const AddEpisode = (props) => {
  const [message, setMessage] = useState(null);
  const [_, modalDispatch] = useContext(ModalContext);
  const [form, setForm] = useState({
    thumbnail: '',
    title: '',
    video_link: '',
    movie_id: '',
  });

  let { data: movies } = useQuery('movies', async () => {
    const response = await API.get('/movies');
    return response.data.data;
  });

  const handleOnChange = (e) => {
    setForm({
      ...form,
      [e.target.name]: e.target.type === 'file' ? e.target.files : e.target.value,
    });
  };

  const handleOnSubmit = useMutation(async (e) => {
    try {
      e.preventDefault();

      const config = {
        headers: {
          'Content-type': 'multipart/form-data',
        },
      };
      const formData = new FormData();
      formData.set('thumbnail', form.thumbnail[0], form.thumbnail[0].name);
      formData.set('title', form.title);
      formData.set('video_link', form.video_link);
      formData.set('movie_id', form.movie_id);

      const response = await API.post('/episode', formData, config);
      console.log('add episode success', response);
    } catch (err) {
      console.log('add episode failed', err);
    }
  });

  return (
    <div className={`${props.className}`}>
      <div className="flex justify-between items-center">
        <h2 className="font-semibold mb-5 text-2xl text-white">Add Episode</h2>
      </div>
      {message && message}
      <form onSubmit={(e) => handleOnSubmit.mutate(e)}>
        <div className="flex gap-x-3">
          <input onChange={handleOnChange} className="w-3/4 mb-3 rounded-md p-2 placeholder-white border-2 border-white bg-zinc-500 focus:outline-none" type="text" name="title" placeholder="Title" />
          <input onChange={handleOnChange} className="w-1/4 mb-3 rounded-md p-2 placeholder-white border-2 border-white bg-zinc-500 focus:outline-none file-input-ghost file:text-white" type="file" name="thumbnail" id="" />
        </div>
        <select onChange={handleOnChange} value={form.movie_id} name="movie_id" className="mb-3 border border-white rounded-md py-1 px-3 bg-zinc-500 w-full">
          <option className="hidden" value={'default'}>
            Movies
          </option>
          {movies?.map((index, id) => (
            <option key={id} value={index?.id}>
              {index?.title}
            </option>
          ))}
        </select>
        <input onChange={handleOnChange} className="w-full mb-3 rounded-md p-2 placeholder-white border-2 border-white bg-zinc-500 focus:outline-none" type="text" name="video_link" id="" placeholder="Link Video" />
        <button className="w-full bg-white py-2 rounded-md mb-2 text-red-700 font-semibold">Save</button>
      </form>
    </div>
  );
};

export default AddEpisode;
