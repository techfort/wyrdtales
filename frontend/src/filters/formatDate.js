import moment from 'moment';

function formatDate(value) {
  return moment(String(value)).format('dddd [the] Do [of] MMMM, YYYY [at] HH:mm');
}

export default formatDate;