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

        public async Task<int?> CreateAsync(PersonRequest person)
        {
            var entity = person.ToPerson();
            await _repository.CreateAsync(entity);
            await _repository.SaveAsync();

            return entity?.Id;
        }

        public async Task DeleteAsync(int id)
        {
            await _repository.DeleteAsync(id);
            await _repository.SaveAsync();
        }

        public async Task<PersonResponse?> GetAsync(int id)
        {
            var ans = await _repository.GetAsync(id);

            return ans is not null ? new PersonResponse(ans) : null;
        }

        public async Task<IEnumerable<PersonResponse>> GetAllAsync()
        {
            return (await _repository.GetAllAsync()).Select(person => new PersonResponse(person));
        }

        public async Task<PersonResponse?> UpdateAsync(int id, PersonRequest person)
        {
            var ans = _repository.GetAsync(id).Result;

            if (ans is null)
                return null;

            ans.Name = person.Name;
            ans.Age = person.Age ?? ans.Age;
            ans.Adress = person.Adress ?? ans.Adress;
            ans.Work = person.Work ?? ans.Work;

            _repository.Update(ans);
            await _repository.SaveAsync();

            return new PersonResponse(ans);
        }
    }
}
