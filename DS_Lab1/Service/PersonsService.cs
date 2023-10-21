using DS_Lab1.CommunicationElements;
using DS_Lab1.Repository;

namespace DS_Lab1.Service
{
    public class PersonsService : IPersonsService
    {
        readonly IPersonsRepository _repository;

        public PersonsService(IPersonsRepository repository)
        {
            _repository = repository;
        }

        public int Create(PersonRequest person)
        {
            var entity = person.ToPerson();
            _repository.CreateAsync(entity);
            _repository.SaveAsync();

            return entity?.Id ?? throw new Exception("Can\'t create person");
        }

        public void Delete(int id)
        {
            _repository.DeleteAsync(id);
        }

        public PersonResponse Get(int id)
        {
            var ans = _repository.GetAsync(id).Result;

            return ans is null ? throw new Exception($"Can\'t find person by id = {id}") : new PersonResponse(ans);
        }

        public IEnumerable<PersonResponse> GetAll()
        {
            return _repository.GetAllAsync().Result.Select(person => new PersonResponse(person));
        }

        public PersonResponse Update(int id, PersonRequest person)
        {
            var ans = _repository.GetAsync(id).Result ?? throw new Exception($"Person with id = {id} does not exist");
            ans.Name = person.Name;
            ans.Age = person.Age ?? ans.Age;
            ans.Adress = person.Adress ?? ans.Adress;
            ans.Work = person.Work ?? ans.Work;

            _repository.Update(ans);

            return new PersonResponse(ans);
        }
    }
}
