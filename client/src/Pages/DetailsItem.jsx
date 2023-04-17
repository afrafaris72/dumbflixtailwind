import React, { useContext, useEffect, useState } from 'react';
import useFetch from '../Config/useFetch';
import { useParams } from 'react-router-dom';
import ReactPlayer from 'react-player';
import ApiConfig from '../Config/ApiConfig';
import { UserContext } from '../Context/UserContext';
import { ModalContext } from '../Context/ModalContext';
import { useQuery } from 'react-query';
import { API } from '../Config/Api';
import placeholderThumb from '../assets/img/thumbnail_video_placeholder.jpg';

const DetailsItem = (props) => {
  const { id } = useParams();
  const { data } = useFetch(`${props.endpoint}${id}`);
  const { tmdb_originalImage, tmdb_w500Image } = ApiConfig;
  const [userState, userDispatch] = useContext(UserContext);
  const [modalState, modalDispatch] = useContext(ModalContext);
  const [urlVideo, setUrlVideo] = useState(null);

  const { data: movie } = useQuery('moviesCache', async () => {
    const response = await API.get(`${props.endpoint}${id}`);
    return response.data.data;
  });

  const { data: episode } = useQuery('episodeCache', async () => {
    const response = await API.get(`/movie/${id}/episodes`);
    return response.data.data;
  });

  useEffect(() => {
    console.log(episode.map((index) => index.thumbnail));
  }, [episode]);

  return (
    <React.Fragment>
      <div className="pt-10 bg-black/60">
        <div className="relative h-[500px]">
          <ReactPlayer
            className={'absolute top-0 left-0'}
            width={'100%'}
            height={'100%'}
            light={
              <div className="px-40 h-[500px]">
                <img className="w-full h-[500px] mx-auto" src={episode ? `${episode.thumbnail}` : placeholderThumb} />
              </div>
            }
            controls={true}
            url={`${urlVideo ? 'https://www.youtube.com/watch?v=${index.key}' : null}`}
          />
        </div>
      </div>

      {userState.user.role === 'admin' ? (
        <div className="bg-black text-end container mx-auto px-8 py-5">
          <button onClick={() => modalDispatch({ type: 'ADD_EPISODE_MODAL' })} className="bg-red-700 text-white px-8 py-2 rounded-md">
            Add Episode
          </button>
        </div>
      ) : null}

      <div className="bg-black">
        <div className={`flex container justify-between mx-auto lg:px-8 ${userState.user.role === 'admin' ? `pb-20` : `py-20`}`}>
          <div className="w-1/2 flex gap-x-8">
            <div className="w-1/3">
              <img src={movie?.thumbnail} alt="" />
            </div>
            <div className="w-2/3 flex flex-col justify-evenly">
              <h2 className="text-3xl font-semibold text-white">{movie?.title}</h2>
              <div className="flex items-center gap-3">
                <p>{movie?.year}</p>
                <p className="border rounded-md p-1">{movie?.category?.name}</p>
              </div>
              <p className="text-justify">{movie?.description}</p>
            </div>
          </div>
          <div className="w-1/2 mx-auto pl-28">
            <img className="rounded-md object-cover w-full h-[330px]" src={`${urlVideo != null ? null : placeholderThumb}`} alt="" />
            <h3>{movie?.title}</h3>
          </div>
        </div>
      </div>
    </React.Fragment>
  );
};

export default DetailsItem;
