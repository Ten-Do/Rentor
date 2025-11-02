import { useParams } from 'react-router-dom';

export const Advertisement = () => {
  const { id } = useParams();
  return <>Advertisement {id}</>;
};

