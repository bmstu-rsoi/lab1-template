using DS_Lab1.Models;
using Microsoft.EntityFrameworkCore;

namespace DS_Lab1.Repository
{
    public class PersonsRepository : IPersonsRepository
    {
        readonly PersonsDbContext _context;

        public PersonsRepository(PersonsDbContext context) 
        {
            _context = context;
        }

        public async Task CreateAsync(Person person)
        {
            await _context.Persons.AddAsync(person);
        }

        public async Task DeleteAsync(int id)
        {
            var item = await GetAsync(id);

            if (item is not null)
                _context.Persons.Remove(item);
        }

        public async Task<IEnumerable<Person>> GetAllAsync()
        {
            return await _context.Persons.ToListAsync();
        }

        public async Task<Person?> GetAsync(int id)
        {
            return await _context.Persons.FirstOrDefaultAsync(e => e.Id.Equals(id));
        }

        public async Task SaveAsync()
        {
            await _context.SaveChangesAsync();
        }

        public void Update(Person person)
        {
            _context.Persons.Update(person);
        }
    }
}
