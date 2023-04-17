import React, { useEffect, useState } from 'react';
import { useMutation, useQuery } from 'react-query';
import { API } from '../Config/Api';

const AdminForm = () => {
  const [preview, setPreview] = useState(null);
  const [form, setForm] = useState({
    title: '',
    year: '',
    thumbnail: '',
    category_id: '',
    description: '',
  });

  const handleOnChange = (e) => {
    setForm({
      ...form,
      [e.target.name]: e.target.type === 'file' ? e.target.files : e.target.value,
    });
  };

  let { data: categories, refetch } = useQuery('categoriesCache', async () => {
    const response = await API.get('/category');
    return response.data.data;
  });

  const handleOnSubmit = useMutation(async (e) => {
    try {
      e.preventDefault();

      console.log(form);
      const config = {
        headers: {
          'Content-type': 'multipart/form-data',
        },
      };

      const formData = new FormData();
      formData.set('thumbnail', form.thumbnail[0], form.thumbnail[0].name);
      formData.set('title', form.title);
      formData.set('year', form.year);
      formData.set('description', form.description);
      formData.set('category_id', form.category_id);

      const response = await API.post('/movie', formData, config);
      console.log('add movie success', response);
    } catch (err) {
      e.preventDefault();
      console.log('add movie failed', err);
      console.log(form);
    }
  });

  return (
    <div className="bg-black container mx-auto pt-28  px-44 h-[100vh]">
      <h2 className="font-bold text-lg text-white mb-5">Add Movie</h2>
      <form onSubmit={(e) => handleOnSubmit.mutate(e)}>
        <div className=" flex gap-3">
          <input onChange={handleOnChange} value={form.title} name="title" type="text" className="mb-3 border border-white rounded-md py-1 px-3 bg-zinc-700 w-3/4" placeholder="Title" />
          <input onChange={handleOnChange} name="thumbnail" type="file" className="mb-3 border border-white rounded-md file:h-full bg-zinc-700 w-1/4" />
        </div>
        <input onChange={handleOnChange} value={form.year} name="year" type="text" className="mb-3 border border-white rounded-md py-1 px-3 bg-zinc-700 w-full" placeholder="Year" />
        <select onChange={handleOnChange} value={form.category_id} name="category_id" className="mb-3 border border-white rounded-md py-1 px-3 bg-zinc-700 w-full">
          <option className="hidden" value={'default'}>
            Category
          </option>
          {categories?.map((index, id) => (
            <option key={id} value={index?.id}>
              {index?.name}
            </option>
          ))}
        </select>
        <textarea value={form.description} onChange={handleOnChange} name="description" className="mb-3 border border-white rounded-md py-1 px-3 bg-zinc-700 w-full" placeholder="Description"></textarea>

        <div className="text-end">
          <button className="bg-red-700 text-white px-12 py-2 rounded-md mt-5">Save</button>
        </div>
      </form>
    </div>
  );
};

export default AdminForm;
